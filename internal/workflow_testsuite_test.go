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

package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetMemoOnStart(t *testing.T) {
	testSuite := &WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	memo := map[string]interface{}{
		"key": make(chan int),
	}
	err := env.SetMemoOnStart(memo)
	require.Error(t, err)

	memo = map[string]interface{}{
		"memoKey": "memo",
	}
	require.Nil(t, env.impl.workflowInfo.Memo)
	err = env.SetMemoOnStart(memo)
	require.NoError(t, err)
	require.NotNil(t, env.impl.workflowInfo.Memo)
}

func TestSetSearchAttributesOnStart(t *testing.T) {
	testSuite := &WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	invalidSearchAttr := map[string]interface{}{
		"key": make(chan int),
	}
	err := env.SetSearchAttributesOnStart(invalidSearchAttr)
	require.Error(t, err)

	searchAttr := map[string]interface{}{
		"CustomIntField": 1,
	}
	err = env.SetSearchAttributesOnStart(searchAttr)
	require.NoError(t, err)
	require.NotNil(t, env.impl.workflowInfo.SearchAttributes)
}
