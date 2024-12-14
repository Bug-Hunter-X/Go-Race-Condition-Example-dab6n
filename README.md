# Go Race Condition Example

This repository demonstrates a simple race condition in Go.  The program uses multiple goroutines to increment a shared counter. Without proper synchronization mechanisms, the final count will likely be incorrect.

## Problem Description

The `bug.go` file contains code that exhibits a race condition. Multiple goroutines concurrently access and modify the `count` variable.  This leads to unpredictable results, as the increments may be lost or overwritten.

## Solution

The `bugSolution.go` file demonstrates how to correctly handle the concurrency issue using a `sync.Mutex` to protect the shared resource.  This ensures that only one goroutine can access and modify the `count` variable at any given time.