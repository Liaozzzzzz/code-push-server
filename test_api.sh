#!/bin/bash

# API测试脚本
BASE_URL="http://localhost:8040"

echo "🧪 开始API测试..."
echo "================================"

# 1. 健康检查
echo "1. 健康检查"
curl -s -X GET "$BASE_URL/health" | jq '.'
echo ""

# 2. 创建用户
echo "2. 创建用户"
curl -s -X POST "$BASE_URL/api/users" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123",
    "role": "developer"
  }' | jq '.'
echo ""

# 3. 测试创建重复用户（应该返回业务错误码1002）
echo "3. 测试创建重复用户"
curl -s -X POST "$BASE_URL/api/users" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123",
    "role": "developer"
  }' | jq '.'
echo ""

# 4. 用户登录
echo "4. 用户登录"
curl -s -X POST "$BASE_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }' | jq '.'
echo ""

# 5. 测试错误登录（应该返回业务错误码1005）
echo "5. 测试错误登录"
curl -s -X POST "$BASE_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "wrongpassword"
  }' | jq '.'
echo ""

# 6. 获取用户列表
echo "6. 获取用户列表"
curl -s -X GET "$BASE_URL/api/users?page=1&size=10" | jq '.'
echo ""

# 7. 获取用户详情
echo "7. 获取用户详情"
curl -s -X GET "$BASE_URL/api/users/1" | jq '.'
echo ""

# 8. 获取不存在的用户（应该返回业务错误码1003）
echo "8. 获取不存在的用户"
curl -s -X GET "$BASE_URL/api/users/999" | jq '.'
echo ""

# 9. 创建应用（需要认证）
echo "9. 创建应用（需要认证）"
curl -s -X POST "$BASE_URL/api/apps" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer test-token" \
  -d '{
    "name": "Test App",
    "platform": "ios",
    "bundle_id": "com.example.testapp",
    "description": "这是一个测试应用"
  }' | jq '.'
echo ""

# 10. 测试无认证创建应用（应该返回HTTP 401）
echo "10. 测试无认证创建应用"
curl -s -X POST "$BASE_URL/api/apps" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test App 2",
    "platform": "android",
    "bundle_id": "com.example.testapp2",
    "description": "这是另一个测试应用"
  }' | jq '.'
echo ""

# 11. 获取应用列表
echo "11. 获取应用列表"
curl -s -X GET "$BASE_URL/api/apps?page=1&size=10" | jq '.'
echo ""

# 12. 获取用户的应用列表（需要认证）
echo "12. 获取用户的应用列表（需要认证）"
curl -s -X GET "$BASE_URL/api/my/apps?page=1&size=10" \
  -H "Authorization: Bearer test-token" | jq '.'
echo ""

# 13. 测试参数错误（应该返回业务错误码1001）
echo "13. 测试参数错误"
curl -s -X POST "$BASE_URL/api/users" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "",
    "email": "invalid-email",
    "password": "123"
  }' | jq '.'
echo ""

echo "✅ API测试完成！"
echo ""
echo "📊 测试结果说明："
echo "- HTTP 200 + code 0：业务成功"
echo "- HTTP 200 + code 1xxx：业务失败（客户端错误）"
echo "- HTTP 200 + code 2xxx：业务失败（服务器错误）"
echo "- HTTP 401：认证失败"
echo "- HTTP 500：系统异常" 