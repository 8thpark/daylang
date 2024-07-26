import React from "react"
import { createBrowserRouter, RouterProvider } from "react-router-dom"
import Dashboard from "./pages/dashboard"
import Learn from "./pages/learn"
import Settings from "./pages/settings"
import Error from "./pages/error"

const router = createBrowserRouter([
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
])

function App() {
  return (
    <React.StrictMode>
      <RouterProvider router={router} />
    </React.StrictMode>
  )
}

export default App
