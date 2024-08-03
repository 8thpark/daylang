import React, { useEffect, useState } from "react"
import { createBrowserRouter, RouterProvider } from "react-router-dom"
import axios from "axios"
import Dashboard from "./pages/dashboard"
import Error from "./pages/error"
import Learn from "./pages/learn"
import Login from "./pages/login"
import Settings from "./pages/settings"
import SelfContext from "./self_context"

function App() {
  const [self, setSelf] = useState(null)
  const [router, setRouter] = useState(null)

  const createRouter = (selfData) => {
    const routes = [
      {
        path: "/",
        element: <Dashboard />,
        errorElement: <Error />,
      },
      {
        path: "/learn",
        element: <Learn />,
      },
      {
        path: "/settings",
        element: <Settings />,
      },
    ]

    if (selfData.mode === "standalone") {
      routes.push({
        path: "/login",
        element: <Login />,
      })
    }

    setRouter(createBrowserRouter(routes))
  }

  const fetchSelfData = async () => {
    try {
      const response = await axios.get("/api/self")
      setSelf(response.data)
      createRouter(response.data)
    } catch (err) {
      console.error("error fetching self data:", err)
    }
  }

  useEffect(() => {
    fetchSelfData()
  }, [])

  if (!router) {
    return null
  }

  return (
    <React.StrictMode>
      <SelfContext.Provider value={self}>
        <RouterProvider router={router} />
      </SelfContext.Provider>
    </React.StrictMode>
  )
}

export default App
