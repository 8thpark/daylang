import React from "react"
import PropTypes from "prop-types"
import { Link } from "react-router-dom"

function Container({ children }) {
  return (
    <div className="mx-auto max-w-2xl px-4 pb-8 pt-24">
      <h1 className="mb-2 text-2xl font-bold">
        <Link to="/">DayLang</Link>
      </h1>

      {children}

      <p>
        <Link to="/logout">Logout</Link>
      </p>
      <p>
        <a href="https://daylang.com/account">Account</a>
      </p>
    </div>
  )
}

Container.propTypes = {
  children: PropTypes.node.isRequired,
}

export default Container
