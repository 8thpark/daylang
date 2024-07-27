import React from "react"
import { Link } from "react-router-dom"

function Navbar() {
  return (
    <div className="navbar bg-base-100">
      <div className="flex-1">
        <Link className="btn btn-ghost -ml-6 text-xl" to="/">
          <img src="/sunset.png" alt="DayLang Icon" className="size-6" />
          DayLang
        </Link>
      </div>
      <div className="flex-none">
        <ul className="menu menu-horizontal px-1">
          <li>
            <a href="https://daylang.com/account">Account</a>
          </li>
          <li>
            <a href="/logout">Logout</a>
          </li>
        </ul>
      </div>
    </div>
  )
}

Navbar.propTypes = {}

export default Navbar
