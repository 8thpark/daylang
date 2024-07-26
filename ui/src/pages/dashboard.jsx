import React from "react"
import { Link } from "react-router-dom"
import Container from "../components/container"

function Dashboard() {
  return (
    <Container>
      <p>Daily vocabulary builder for language learners.</p>
      <p>
        <Link to="/learn">Learn</Link>
      </p>
      <p>
        <Link to="/settings">Settings</Link>
      </p>
    </Container>
  )
}

export default Dashboard
