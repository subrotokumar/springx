#!/bin/bash
# Example usage of SpringX CLI

echo "=== SpringX CLI Examples ==="
echo ""

echo "1. List available dependencies"
echo "   Command: springx init --list"
echo ""

echo "2. Generate a basic web application"
echo "   Command: springx init my-web-app"
echo ""

echo "3. Generate with Gradle and Java 21"
echo "   Command: springx init modern-app --build=gradle --java=21"
echo ""

echo "4. Generate a REST API with multiple dependencies"
echo "   Command: springx init rest-api --deps=web,data-jpa,validation,actuator"
echo ""

echo "5. Generate with custom group and package"
echo "   Command: springx init my-service --group=io.mycompany --package=io.mycompany.myservice"
echo ""

echo "6. Generate a secure web application"
echo "   Command: springx init secure-app --deps=web,security,thymeleaf"
echo ""

echo "7. Generate to a specific directory"
echo "   Command: springx init my-app --output=/tmp/spring-projects/my-app"
echo ""

echo "=== Try it yourself! ==="
echo "Run: springx init --help for more information"
