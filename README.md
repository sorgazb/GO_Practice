# GO_Practice

![Go](https://img.shields.io/badge/Go-1.21+-00add8?style=for-the-badge&logo=go&logoColor=white)&nbsp;![Aprendizaje](https://img.shields.io/badge/Aprendizaje-Autodidacta-f97316?style=for-the-badge)&nbsp;![Estado](https://img.shields.io/badge/Estado-En%20progreso-22c55e?style=for-the-badge)

> **GO_Practice** es un repositorio de ejercicios y ejemplos de introducción al lenguaje **Go (Golang)**, organizado por temas progresivos: desde la sintaxis básica hasta funciones, estructuras y agrupación de datos.

---

## 📋 Contenido del Repositorio

El repositorio está dividido en unidades temáticas numeradas de menor a mayor complejidad:

### 01 · Introducción

Primeros pasos con Go: sintaxis básica, declaración de variables y tipos primitivos.

- `hola_mundo.go` — El clásico primer programa en Go.
- `variables.go` — Declaración de variables con `var`.
- `declaracion_corta.go` — Operador `:=` y declaración corta.
- `tipos.go` — Tipos de datos básicos en Go.
- `valor_cero.go` — Valor cero de cada tipo de dato.
- `nombre.go` — Entrada y salida por consola.
- `ejercicios_01/` — Ejercicios prácticos de la unidad.

### 02 · Tipos de Dato

Estudio en profundidad de los tipos de dato disponibles en Go: enteros, flotantes, booleanos, strings y conversiones de tipo.

### 03 · Control de Flujo

Estructuras de control: condicionales `if/else`, bucles `for`, `switch` y uso de `defer`.

### 04 · Agrupación de Datos

Trabajar con colecciones: arrays, slices y maps en Go.

### 05 · Estructuras

Definición de `structs`, métodos asociados y composición de tipos.

### 06 · Funciones

Funciones con múltiples valores de retorno, funciones variadic, closures y funciones como valores.

---

## 🏗️ Estructura del Proyecto

```txt
GO_Practice/
├── 01.introduccion/
│   ├── hola_mundo.go
│   ├── variables.go
│   ├── declaracion_corta.go
│   ├── tipos.go
│   ├── valor_cero.go
│   ├── nombre.go
│   └── ejercicios_01/
├── 02.tipos_dato/
├── 03.control_flujo/
├── 04.agrupacion_datos/
├── 05.estructuras/
├── 06.funciones/
└── README.md
```

---

## ⚙️ Requisitos y Ejecución

Clonar el repositorio:
```bash
git clone https://github.com/sorgazb/GO_Practice.git
cd GO_Practice
```

Asegurarse de tener Go instalado ([golang.org/dl](https://golang.org/dl)):
```bash
go version
```

Ejecutar cualquier archivo directamente:
```bash
go run 01.introduccion/hola_mundo.go
```

O compilar y ejecutar:
```bash
go build 01.introduccion/hola_mundo.go
./hola_mundo
```

---

## 🤝 Contribución

Haz fork del repositorio.

Crea una rama de trabajo:
```bash
git checkout -b feature/mi-nueva-funcionalidad
```

Realiza tus cambios y haz commit.

Abre un Pull Request describiendo tus mejoras.

---

<p align="center">Aprendizaje autodidacta de Go &nbsp;·&nbsp; Sergio Orgaz Bravo</p>
