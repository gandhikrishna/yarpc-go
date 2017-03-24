// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package pally

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGauge(t *testing.T) {
	r := NewRegistry(Labeled(Labels{"service": "users"}))
	gauge, err := r.NewGauge(Opts{
		Name:        "test_gauge",
		Help:        "Some help.",
		ConstLabels: Labels{"foo": "bar"},
	})
	require.NoError(t, err, "Unexpected error constructing gauge.")

	scope := newTestScope()
	stop, err := r.Push(scope, _tick)
	require.NoError(t, err, "Unexpected error starting Tally push.")

	gauge.Inc()
	gauge.Store(42)
	assert.Equal(t, int64(42), gauge.Load(), "Unexpected in-memory gauge value.")

	time.Sleep(5 * _tick)
	gauge.Store(4)
	assert.Equal(t, int64(4), gauge.Load(), "Unexpected in-memory gauge value after sleep.")

	stop()

	export := TallyExpectation{
		Type:   "gauge",
		Name:   "test_gauge",
		Labels: Labels{"foo": "bar", "service": "users"},
		Value:  4,
	}
	export.Test(t, scope)

	assertPrometheusText(t, r, "# HELP test_gauge Some help.\n"+
		"# TYPE test_gauge gauge\n"+
		`test_gauge{foo="bar",service="users"} 4`)
}

func TestGaugeVec(t *testing.T) {
	r := NewRegistry(Labeled(Labels{"service": "users"}))
	vec, err := r.NewGaugeVector(Opts{
		Name:           "test_gauge",
		Help:           "Some help.",
		ConstLabels:    Labels{"foo": "bar"},
		VariableLabels: []string{"baz"},
	})
	require.NoError(t, err, "Unexpected error constructing vector.")

	scope := newTestScope()
	stop, err := r.Push(scope, _tick)
	require.NoError(t, err, "Unexpected error starting Tally push.")

	gauge, err := vec.Get("a")
	require.NoError(t, err, "Unexpected error getting a gauge with correct number of labels.")

	gauge.Store(11)
	assert.Equal(t, int64(11), gauge.Load(), "Unexpected in-memory gauge value.")

	vec.MustGet("a").Add(2)
	vec.MustGet("a").Inc()
	assert.Equal(t, int64(14), gauge.Load(), "Unexpected in-memory gauge value.")

	time.Sleep(5 * _tick)
	gauge.Store(5)
	assert.Equal(t, int64(5), gauge.Load(), "Unexpected in-memory gauge value after sleep.")

	stop()

	export := TallyExpectation{
		Type:   "gauge",
		Name:   "test_gauge",
		Labels: Labels{"foo": "bar", "service": "users", "baz": "a"},
		Value:  5,
	}
	export.Test(t, scope)

	assertPrometheusText(t, r, "# HELP test_gauge Some help.\n"+
		"# TYPE test_gauge gauge\n"+
		`test_gauge{baz="a",foo="bar",service="users"} 5`)
}

func TestGaugeVecInvalidLabelValues(t *testing.T) {
	r := NewRegistry()
	vec, err := r.NewGaugeVector(Opts{
		Name:           "test_counter",
		Help:           "Some help.",
		VariableLabels: []string{"foo"},
	})
	require.NoError(t, err, "Unexpected error constructing vector.")

	_, err = vec.Get("foo:")
	assert.Error(t, err, "Expected an error using invalid label values.")
	assert.Panics(t, func() { vec.MustGet("foo:") }, "Expected an error using invalid label values.")

	_, err = vec.Get("bar", "baz")
	assert.Error(t, err, "Expected a panic using too many label values.")
	assert.Panics(t, func() { vec.MustGet("bar", "baz") }, "Expected a panic using too many label values.")
}