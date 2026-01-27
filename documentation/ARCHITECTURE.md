# AegisFlux Architecture

Este documento describe la arquitectura del proyecto AegisFlux.

## 1. Overview
El proyecto está dividido en varios módulos: agent, api, honeypot, jenkins, k8s, etc.

## 2. GO Module
- /bin/: ejecutables
- /build/: archivos de construcción
- /dist/: distribución
- /out/: salida

## 3. Angular Module
- /dashboard/: frontend Angular
- node_modules/: dependencias
- dist/: build del frontend
- .angular/, .cache/: cache del framework

## 4. Docker
- docker/*.log
- docker-compose.override.yml
- .dockerignore

## 5. Kubernetes
- *.yaml~ y archivos sensibles

## 6. VS Code
- Configuración local

## 7. Logs, Temp, Backups
- *.log, *.tmp, *.swp, *.bak

## 8. Python / Scripts
- __pycache__/
- *.pyc

## 9. Honeypot Logs
- /honeypot/*.log

## 10. Environment
- .env, .env.*
