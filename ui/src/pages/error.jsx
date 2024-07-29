import React from "react"
import { Link, useRouteError } from "react-router-dom"
import { useDocumentTitle } from "@uidotdev/usehooks"
import Container from "../components/container"

function Error() {
  useDocumentTitle("DayLang — Error")

  const error = useRouteError()
  console.error(error) // eslint-disable-line no-console

  return (
    <Container>
      <div className="prose">
        <p>Oops! Sorry, an unexpected error has occurred.</p>
        <p>
          <i>{error.statusText || error.message}</i>
        </p>

        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
        </ul>
      </div>
    </Container>
  )
}

export default Error
