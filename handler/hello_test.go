package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloWorld(t *testing.T) {
	// 创建测试上下文
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 调用处理函数
	HelloWorld(c)

	// 验证状态码
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// 验证响应内容
	expected := "Hello World!"
	if !strings.Contains(w.Body.String(), expected) {
		t.Errorf("expected response body to contain %q, got %q", expected, w.Body.String())
	}
}

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Ping(c)

	// 验证状态码
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// 验证响应内容
	expected := "pong"
	if !strings.Contains(w.Body.String(), expected) {
		t.Errorf("expected response body to contain %q, got %q", expected, w.Body.String())
	}
}
