import React from "react"
import { Link, useRouteError } from "react-router-dom"
import Container from "../components/container"

function Error() {
  const error = useRouteError()
  console.error(error) // eslint-disable-line no-console

  return (
    <Container>
      <div className="prose">
        <h1>Oops!</h1>

        <p>Sorry, an unexpected error has occurred.</p>
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
