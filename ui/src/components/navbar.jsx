import React, { useContext } from "react"
import { Link, useNavigate } from "react-router-dom"
import SelfContext from "../self_context"

function Navbar() {
  const self = useContext(SelfContext)
  const navigate = useNavigate()

  const handleLogout = () => {
    if (self && self.mode === "bridge") {
      window.location.href = `${self.billingRedirect}/logout`
      return
    }

    navigate("/login")
  }

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
          {self && self.mode === "bridge" && (
            <li>
              <a href={`${self.billingRedirect}/user/settings`}>Account</a>
            </li>
          )}

          <li>
            <button type="button" onClick={handleLogout}>
              Logout
            </button>
          </li>
        </ul>
      </div>
    </div>
  )
}

Navbar.propTypes = {}

export default Navbar
