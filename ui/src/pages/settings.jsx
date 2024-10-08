import React from "react"
import { useDocumentTitle } from "@uidotdev/usehooks"
import Container from "../components/container"

function Settings() {
  useDocumentTitle("DayLang — Settings")
  return (
    <Container>
      <div className="prose">
        <p>This is the settings page.</p>
      </div>
    </Container>
  )
}

export default Settings
