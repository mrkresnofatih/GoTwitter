import React from 'react'
import NavbarItem from './NavbarItem'

const Navbar = () => {
  return (
    <>
        <NavbarItem text={"Home"} href={"/"} />
        <NavbarItem text={"Profile"} href={"/"} />
        <NavbarItem text={"Explore"} href={"/"} />
        <NavbarItem text={"Bookmarks"} href={"/"} />
        <NavbarItem text={"Settings"} href={"/"} />
        <NavbarItem text={"Logout"} href={"/"} />
    </>
  )
}

export default Navbar