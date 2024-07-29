import React from "react"
import { Link } from "react-router-dom"
import { useDocumentTitle } from "@uidotdev/usehooks"
import Container from "../components/container"

function Learn() {
  useDocumentTitle("DayLang — Learn")
  return (
    <Container>
      <div className="prose">
        <p>This is the learn page.</p>

        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
        </ul>
      </div>
    </Container>
  )
}

export default Learn
