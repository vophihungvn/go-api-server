# Sample API server written in Go

## Libraries

- Web framework: [Gin](https://github.com/gin-gonic/gin)
- ORM: [Gorm](https://gorm.io/)
- Migration: [gormigrate.v1](https://godoc.org/gopkg.in/gormigrate.v1)

## Project structure

| Folder       | Purpose                                   |
| ------------ | ----------------------------------------- |
| src/routers  | Setup routing, mapping route with handler |
| src/services | Service code, doing some specific funtion |
| src/models   | Define models, migrations,...             |
