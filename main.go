package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stack, needDB := promptForDetails()
	createProjectStructure(stack, needDB)
}

// Prompt user for project details
func promptForDetails() (string, bool) {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for stack selection
	fmt.Println("Choose your stack:")
	fmt.Println("1. React Vite + Express")
	fmt.Println("2. Next.js")
	fmt.Print("Enter your choice (1/2): ")
	stackChoice, _ := reader.ReadString('\n')
	stackChoice = strings.TrimSpace(stackChoice)

	var stack string
	if stackChoice == "1" {
		stack = "React Vite + Express"
	} else if stackChoice == "2" {
		stack = "Next.js"
	} else {
		fmt.Println("Invalid choice, defaulting to React Vite + Express.")
		stack = "React Vite + Express"
	}

	// Prompt for database confirmation
	fmt.Print("Do you need a database? (y/n): ")
	needDBChoice, _ := reader.ReadString('\n')
	needDBChoice = strings.TrimSpace(strings.ToLower(needDBChoice))

	var needDB bool
	if needDBChoice == "y" || needDBChoice == "yes" {
		needDB = true
	} else {
		needDB = false
	}

	return stack, needDB
}

// Create the project structure based on user input
func createProjectStructure(stack string, needDB bool) {
	fmt.Println("Project will use:", stack)
	if needDB {
		fmt.Println("Setting up database...")
		setupDatabase()
		createDockerFiles()
	} else {
		fmt.Println("No database setup required.")
	}

	// Create frontend and backend structure based on stack
	if stack == "React Vite + Express" {
		createReactViteTemplate()
		createExpressTemplate()
	} else if stack == "Next.js" {
		createNextJSTemplate()
	}
	fmt.Println("Project structure generated!")
}

// Create React Vite TypeScript project structure
func createReactViteTemplate() {
	os.MkdirAll("project/frontend", os.ModePerm)

	// Create package.json for React Vite with TypeScript
	createFile("project/frontend/package.json", `{
  "name": "react-vite-app",
  "version": "1.0.0",
  "scripts": {
    "dev": "vite",
    "build": "vite build",
    "preview": "vite preview"
  },
  "dependencies": {
    "react": "^18.0.0",
    "react-dom": "^18.0.0"
  },
  "devDependencies": {
    "vite": "^3.0.0",
    "typescript": "^4.4.4",
    "@types/react": "^18.0.0",
    "@types/react-dom": "^18.0.0"
  }
}`)

	// Create tsconfig.json for TypeScript
	createFile("project/frontend/tsconfig.json", `{
  "compilerOptions": {
    "target": "ESNext",
    "useDefineForClassFields": true,
    "lib": ["DOM", "DOM.Iterable", "ESNext"],
    "allowJs": false,
    "skipLibCheck": true,
    "esModuleInterop": false,
    "allowSyntheticDefaultImports": true,
    "strict": true,
    "forceConsistentCasingInFileNames": true,
    "module": "ESNext",
    "moduleResolution": "Node",
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "react-jsx"
  },
  "include": ["src"]
}`)

	// Create Vite configuration file
	createFile("project/frontend/vite.config.ts", `import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()]
})
`)

	// Create index.html and basic React files
	os.MkdirAll("project/frontend/src", os.ModePerm)
	createFile("project/frontend/src/main.tsx", `
import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
`)

	createFile("project/frontend/src/App.tsx", `
import { useState } from 'react'
import './App.css'

function App() {
	const [count, setCount] = useState(0)

	return (
		<div className="App">
			<h1>Welcome to React Vite with TypeScript</h1>
			<div>
				<button onClick={() => setCount(count + 1)}>
					count is {count}
				</button>
			</div>
		</div>
	)
}

export default App
`)

	createFile("project/frontend/src/index.css", `
body {
	margin: 0;
	font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}
`)

	createFile("project/frontend/index.html", `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<title>React Vite + TypeScript</title>
</head>
<body>
	<div id="root"></div>
	<script type="module" src="/src/main.tsx"></script>
</body>
</html>
`)
	fmt.Println("React Vite frontend with TypeScript created.")
}

// Create Express TypeScript project structure
func createExpressTemplate() {
	os.MkdirAll("project/backend", os.ModePerm)

	// Create package.json for Express with TypeScript
	createFile("project/backend/package.json", `{
  "name": "express-app",
  "version": "1.0.0",
  "scripts": {
    "start": "ts-node-dev --respawn --transpile-only src/index.ts"
  },
  "dependencies": {
    "express": "^4.0.0",
    "@types/express": "^4.17.0"
  },
  "devDependencies": {
    "typescript": "^4.4.4",
    "ts-node-dev": "^1.1.8"
  }
}`)

	// Create tsconfig.json for TypeScript
	createFile("project/backend/tsconfig.json", `{
  "compilerOptions": {
    "target": "ESNext",
    "module": "CommonJS",
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true
  }
}`)

	// Correct index.ts (Express entry file)
	createFile("project/backend/src/index.ts", `
import express from 'express';

const app = express();
const port = 3001;

app.get('/', (req, res) => {
  res.json({ message: 'Welcome to the Express backend!' });
});

app.listen(port, () => {
  console.log('Server is running on port:' ${port});
});
`)

	fmt.Println("Express backend with TypeScript created.")
}

// Create Next.js TypeScript project structure
func createNextJSTemplate() {
	os.MkdirAll("project/frontend", os.ModePerm)

	// Create package.json for Next.js with TypeScript
	createFile("project/frontend/package.json", `{
  "name": "next-app",
  "version": "1.0.0",
  "scripts": {
    "dev": "next dev",
    "build": "next build",
    "start": "next start"
  },
  "dependencies": {
    "next": "12.0.0",
    "react": "^18.0.0",
    "react-dom": "^18.0.0",
    "typescript": "^4.4.4",
    "@types/react": "^18.0.0",
    "@types/node": "^17.0.0"
  }
}`)
	fmt.Println("Next.js frontend with TypeScript created.")
}

// Set up Prisma and the database
func setupDatabase() {
	// Create prisma folder
	os.MkdirAll("project/backend/prisma", os.ModePerm)

	// Create schema.prisma
	createFile("project/backend/prisma/schema.prisma", `
datasource db {
  provider = "postgresql"
  url      = env("DATABASE_URL")
}

generator client {
  provider = "prisma-client-js"
}

model User {
  id    Int    @id @default(autoincrement())
  name  String
  email String @unique
}
`)

	// Create .env.local
	createFile("project/backend/.env.local", `
DATABASE_URL="postgresql://postgres:postgres@localhost:10001/mydb"
`)

	fmt.Println("Prisma and database setup created.")
}

// Docker Compose and Dockerfiles
func createDockerFiles() {
	// Dockerfile for backend
	createFile("project/backend/Dockerfile", `
FROM node:14
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3001
CMD ["npm", "start"]
`)

	// Dockerfile for frontend
	createFile("project/frontend/Dockerfile", `
FROM node:14
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD ["npm", "run", "dev"]
`)

	// Updated Docker Compose file with PostgreSQL configuration
	createFile("project/docker-compose.yml", `
version: "3.9"
services:
  frontend:
    build: ./frontend
    ports:
      - "3000:3000"
  backend:
    build: ./backend
    ports:
      - "3001:3001"
  postgres1:
    image: postgres:13
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
    command: -c fsync=off -c full_page_writes=off -c synchronous_commit=off -c max_connections=500
    ports:
      - "10001:5432"
`)

	fmt.Println("Docker files created with PostgreSQL configuration.")
}

// Helper function to create files
func createFile(path string, content string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.WriteString(content)
}
