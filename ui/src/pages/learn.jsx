import React from "react"
import { useDocumentTitle } from "@uidotdev/usehooks"
import Container from "../components/container"

function Learn() {
  useDocumentTitle("DayLang — Learn")
  return (
    <Container>
      <div className="prose">
        <p>This is the learn page.</p>
      </div>
    </Container>
  )
}

export default Learn
