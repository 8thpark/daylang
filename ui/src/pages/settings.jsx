import React from "react"
import { Link } from "react-router-dom"
import Container from "../components/container"

function Settings() {
  return (
    <Container>
      <p>This is the settings page.</p>
      <p>
        <Link to="/">Home</Link>
      </p>
    </Container>
  )
}

export default Settings
