import React from "react"
import { Link } from "react-router-dom"
import { useDocumentTitle } from "@uidotdev/usehooks"
import Container from "../components/container"

function Settings() {
  useDocumentTitle("DayLang — Settings")
  return (
    <Container>
      <div className="prose">
        <p>This is the settings page.</p>

        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
        </ul>
      </div>
    </Container>
  )
}

export default Settings
