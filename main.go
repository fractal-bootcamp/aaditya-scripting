package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	stack, needDB := promptForDetails()
	createProjectStructure(stack, needDB)
	runInstallations()
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
	setupAuth() // Set up Firebase authentication
	fmt.Println("Project structure generated!")
}

// Create React Vite TypeScript project structure with Firebase auth
func createReactViteTemplate() {
	os.MkdirAll("project/frontend", os.ModePerm)

	// Create package.json for React Vite with TypeScript and Firebase
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
    "react-dom": "^18.0.0",
    "firebase": "^9.6.1"
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

	// Create basic Firebase authentication files
	createFile("project/frontend/src/firebaseConfig.ts", `
import { initializeApp } from 'firebase/app';
import { getAuth } from 'firebase/auth';

const firebaseConfig = {
  apiKey: "YOUR_API_KEY",
  authDomain: "YOUR_AUTH_DOMAIN",
  projectId: "YOUR_PROJECT_ID",
  storageBucket: "YOUR_STORAGE_BUCKET",
  messagingSenderId: "YOUR_MESSAGING_SENDER_ID",
  appId: "YOUR_APP_ID"
};

const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
`)

	createFile("project/frontend/src/Auth.tsx", `
import React, { useState } from 'react';
import { auth } from './firebaseConfig';
import { signInWithEmailAndPassword, createUserWithEmailAndPassword } from 'firebase/auth';

function Auth() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isLogin, setIsLogin] = useState(true);

  const handleAuth = async () => {
    try {
      if (isLogin) {
        await signInWithEmailAndPassword(auth, email, password);
        alert('Logged in!');
      } else {
        await createUserWithEmailAndPassword(auth, email, password);
        alert('Account created!');
      }
    } catch (error) {
      alert('Error: ' + error.message);
    }
  };

  return (
    <div>
      <h2>{isLogin ? 'Login' : 'Sign Up'}</h2>
      <input type="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={handleAuth}>{isLogin ? 'Login' : 'Sign Up'}</button>
      <button onClick={() => setIsLogin(!isLogin)}>Switch to {isLogin ? 'Sign Up' : 'Login'}</button>
    </div>
  );
}

export default Auth;
`)

	fmt.Println("React Vite frontend with Firebase Authentication created.")
}

// Create Next.js TypeScript project structure
func createNextJSTemplate() {
	os.MkdirAll("project/frontend", os.ModePerm)

	// Create package.json for Next.js with TypeScript and Firebase
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
    "firebase": "^9.6.1",
    "@types/react": "^18.0.0",
    "@types/node": "^17.0.0"
  }
}`)

	// Create tsconfig.json for TypeScript
	createFile("project/frontend/tsconfig.json", `{
  "compilerOptions": {
    "target": "ESNext",
    "module": "ESNext",
    "strict": true,
    "jsx": "preserve",
    "esModuleInterop": true,
    "moduleResolution": "Node",
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "lib": ["DOM", "DOM.Iterable", "ESNext"]
  },
  "include": ["next-env.d.ts", "**/*.ts", "**/*.tsx"],
  "exclude": ["node_modules"]
}`)

	// Create basic Firebase configuration for Next.js
	createFile("project/frontend/src/firebaseConfig.ts", `
import { initializeApp } from 'firebase/app';
import { getAuth } from 'firebase/auth';

const firebaseConfig = {
  apiKey: "YOUR_API_KEY",
  authDomain: "YOUR_AUTH_DOMAIN",
  projectId: "YOUR_PROJECT_ID",
  storageBucket: "YOUR_STORAGE_BUCKET",
  messagingSenderId: "YOUR_MESSAGING_SENDER_ID",
  appId: "YOUR_APP_ID"
};

const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
`)

	// Create index page for Next.js with basic auth integration
	createFile("project/frontend/pages/index.tsx", `
import React, { useState } from 'react';
import { auth } from '../src/firebaseConfig';
import { signInWithEmailAndPassword, createUserWithEmailAndPassword } from 'firebase/auth';

const Home = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [isLogin, setIsLogin] = useState(true);

  const handleAuth = async () => {
    try {
      if (isLogin) {
        await signInWithEmailAndPassword(auth, email, password);
        alert('Logged in!');
      } else {
        await createUserWithEmailAndPassword(auth, email, password);
        alert('Account created!');
      }
    } catch (error) {
      alert('Error: ' + error.message);
    }
  };

  return (
    <div>
      <h2>{isLogin ? 'Login' : 'Sign Up'}</h2>
      <input type="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} />
      <input type="password" placeholder="Password" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button onClick={handleAuth}>{isLogin ? 'Login' : 'Sign Up'}</button>
      <button onClick={() => setIsLogin(!isLogin)}>Switch to {isLogin ? 'Sign Up' : 'Login'}</button>
    </div>
  );
};

export default Home;
`)

	fmt.Println("Next.js frontend with Firebase Authentication created.")
}

// Create Express TypeScript project structure with JWT Auth
func createExpressTemplate() {
	os.MkdirAll("project/backend", os.ModePerm)

	// Create package.json for Express with TypeScript and JWT Auth
	createFile("project/backend/package.json", `{
  "name": "express-app",
  "version": "1.0.0",
  "scripts": {
    "start": "ts-node-dev --respawn --transpile-only src/index.ts"
  },
  "dependencies": {
    "express": "^4.0.0",
    "@types/express": "^4.17.0",
    "jsonwebtoken": "^8.5.1",
    "bcryptjs": "^2.4.3"
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

	// Create index.ts (Express entry file) with JWT Auth
	createFile("project/backend/src/index.ts", `
import express from 'express';
import jwt from 'jsonwebtoken';
import bcrypt from 'bcryptjs';

const app = express();
const port = 3001;
const users = [{ email: 'test@example.com', password: bcrypt.hashSync('password123', 10) }];

app.use(express.json());

app.post('/login', (req, res) => {
  const { email, password } = req.body;
  const user = users.find(u => u.email === email);
  if (user && bcrypt.compareSync(password, user.password)) {
    const token = jwt.sign({ email: user.email }, 'secretKey', { expiresIn: '1h' });
    res.json({ token });
  } else {
    res.status(401).json({ message: 'Invalid credentials' });
  }
});

app.listen(port, () => {
  console.log(\'Server is running on port ${port}\');
});
`)

	fmt.Println("Express backend with JWT Authentication created.")
}

// Set up Firebase Authentication in the frontend
func setupAuth() {
	fmt.Println("Setting up Firebase authentication...")

	// Instructions for user to add Firebase project credentials
	fmt.Println(`
1. Go to https://console.firebase.google.com/ and create a new project.
2. Set up Firebase Authentication in the Firebase console.
3. Replace the placeholder values in src/firebaseConfig.ts with your Firebase project's API key, authDomain, and other details.
`)
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

	// Docker Compose file
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

	fmt.Println("Docker files created.")
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

// Automatically run npm install for frontend and backend
func runInstallations() {
	fmt.Println("Running installations for frontend and backend...")

	// Run npm install in frontend
	frontendCmd := exec.Command("npm", "install")
	frontendCmd.Dir = "project/frontend"
	err := frontendCmd.Run()
	if err != nil {
		fmt.Println("Error installing frontend dependencies:", err)
	} else {
		fmt.Println("Frontend dependencies installed.")
	}

	// Run npm install in backend
	backendCmd := exec.Command("npm", "install")
	backendCmd.Dir = "project/backend"
	err = backendCmd.Run()
	if err != nil {
		fmt.Println("Error installing backend dependencies:", err)
	} else {
		fmt.Println("Backend dependencies installed.")
	}
}
