import React from "react"
import { Link } from "react-router-dom"
import Container from "../components/container"

function Dashboard() {
  return (
    <Container>
      <div className="prose">
        <p>Daily vocabulary builder for language learners.</p>

        <ul>
          <li>
            <Link to="/learn">Learn</Link>
          </li>
          <li>
            <Link to="/settings">Settings</Link>
          </li>
        </ul>
      </div>
    </Container>
  )
}

export default Dashboard
