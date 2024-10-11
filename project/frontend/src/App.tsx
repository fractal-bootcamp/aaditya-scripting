
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
