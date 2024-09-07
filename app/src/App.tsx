import { useState } from "react";
import reactLogo from "@assets/react.svg";
import viteLogo from "/vite.svg";
import Card from "./components/card"

function App() {
  const [count, setCount] = useState(0);

  return (
    <main className="mx-auto container min-h-screen min-w-80 place-items-center justify-center px-4 flex ">
      <section>
        <div>
          <a href="https://vitejs.dev" target="_blank" className="inline-block">
            <img src={viteLogo} className="logo" alt="Vite logo" />
          </a>
          <a href="https://react.dev" target="_blank" className="inline-block">
            <img src={reactLogo} className="logo react" alt="React logo" />
          </a>
        </div>
        <Card title="Go get this shoe" description="Some shoe that that looks amazing" imageUrl="https://images.unsplash.com/photo-1485736231968-0c8ad5c9e174?q=80&w=1470&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D" />
        <h1 className="text-5xl font-bold">Vite + React</h1>
        <div className="p-8 text-center">
          <button
            className="btn inline-block"
            onClick={() => setCount((count) => count + 1)}
          >
            count is {count}
          </button>
          <p>
            Edit <code>src/App.tsx</code> and save to test HMR
          </p>
        </div>
        <p className="text-[#888]">
          Click on the Vite and React logos to learn more
        </p>
      </section>
    </main>
  );
}

export default App;
