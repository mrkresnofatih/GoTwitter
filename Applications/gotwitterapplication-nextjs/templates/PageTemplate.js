import { css } from '@emotion/css'
import Head from 'next/head'
import React from 'react'
import Navbar from '../components/Navbar'

const PageTemplate = ({children}) => {
  return (
    <div className={PageTemplateStyles.Body}>
        <Head>
            <title>Twitter | Home</title>
            <meta name="description" content="Twitter Clone using Golang Echo, MongoDB (Match, Lookup, Aggregation Piplines), & GoAwth OpenIDConnect Identity Provider (GORM, MySQL, NextJs)" />
            <link rel="icon" href="/logo-96.png" />
        </Head>
        <div className={PageTemplateStyles.Inner}>
            <div className={PageTemplateStyles.Side}>
                <div className={PageTemplateStyles.Padding}><Navbar/></div>
            </div>
            <div className={PageTemplateStyles.Core}>{children}</div>
            <div className={PageTemplateStyles.Side}>
                <div className={PageTemplateStyles.Padding}>Trends</div>
            </div>
        </div>
    </div>
  )
}

export default PageTemplate

const PageTemplateStyles = {
    Body: css`
        display: flex;
        flex-direction: column;
        align-items: center;
        height: 100vh;
        width: 100vw;
        background-color: black;
    `,
    Inner: css`
        width: 1000px;
        display: flex;
        height: 100vh;
    `,
    Side: css`
        display: flex;
        flex: 1;
    `,
    Core: css`
        display: flex;
        flex: 2.5;
        /* background-color: coral; */
        border-left: 1px solid rgba(255, 255, 255, 0.1);
        border-right: 1px solid rgba(255, 255, 255, 0.1);
    `,
    Padding: css`
        margin: 18px;
        display: flex;
        flex-direction: column;
        flex: 1;
    `
}