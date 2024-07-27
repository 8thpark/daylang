import React from "react"
import PropTypes from "prop-types"
import Navbar from "./navbar"

function Container({ children }) {
  return (
    <div className="mx-auto max-w-2xl px-4 pb-8 pt-24">
      <Navbar />

      {children}
    </div>
  )
}

Container.propTypes = {
  children: PropTypes.node.isRequired,
}

export default Container
