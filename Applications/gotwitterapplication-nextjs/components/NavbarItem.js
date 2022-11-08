import { css } from '@emotion/css'
import { useRouter } from 'next/router'
import React from 'react'

const NavbarItem = ({text, href}) => {
    const router = useRouter()
    const handlerClick = (href) => (e) => {
        e.preventDefault()
        router.push(href)
    }

    return (
        <label className={navbarItemStyles.body} onClick={handlerClick(href)}>{text}</label>
    )
}

export default NavbarItem

const navbarItemStyles = {
    body: css`
        padding: 12px 18px;
        display: flex;
        flex-direction: column;
        color: white;
        font-size: 18px;
        font-weight: 400;
        border-radius: 50px;

        :hover {
            cursor: pointer;
            background-color: rgba(255, 255, 255, 0.025);
            color: #82CD47;
        }
    `
}