import React from "react"
import { Link } from "react-router-dom"
import Container from "../components/container"

function Learn() {
  return (
    <Container>
      <p>This is the learn page.</p>
      <p>
        <Link to="/">Home</Link>
      </p>
    </Container>
  )
}

export default Learn
