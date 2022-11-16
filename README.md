Go + React Socket Chat
===

## Run & Test

- Backend
  ```
  CD backed
  
  go run main.go
  
  ```
  
- Frontend
  ```
  CD frontend
  
  yarn or npm install
  
  yarn start 
  
  ```


## Directory Hierarchy
```
|—— .DS_Store
|—— backend
|    |—— go.mod
|    |—— go.sum
|    |—— main.go
|    |—— pkg
|        |—— websocket
|            |—— client.go
|            |—— pool.go
|            |—— websocket.go
|—— diagram
|    |—— .DS_Store
|    |—— GO + React Real Time Chat.drawio.png
|—— frontend
|    |—— .gitignore
|    |—— README.md
|    |—— package.json
|    |—— public
|        |—— favicon.ico
|        |—— index.html
|        |—— logo192.png
|        |—— logo512.png
|        |—— manifest.json
|        |—— robots.txt
|    |—— src
|        |—— App.css
|        |—— App.js
|        |—— App.test.js
|        |—— api
|            |—— index.js
|        |—— components
|            |—— ChatHistory
|                |—— ChatHistory.jsx
|                |—— ChatHistory.scss
|                |—— index.js
|            |—— ChatInput
|                |—— ChatInput.jsx
|                |—— ChatInput.scss
|                |—— index.js
|            |—— Header
|                |—— Header.jsx
|                |—— Header.scss
|                |—— index.js
|            |—— Message
|                |—— Message.jsx
|                |—— Message.scss
|                |—— index.js
|        |—— index.css
|        |—— index.js
|        |—— logo.svg
|        |—— reportWebVitals.js
|        |—— setupTests.js
|    |—— yarn.lock
```
