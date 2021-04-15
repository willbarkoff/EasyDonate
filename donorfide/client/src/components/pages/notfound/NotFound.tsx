import * as React from 'react'
import Hero from "../../ui/Hero";
import Page from "../Page";

export default function NotFound() {
    return <Page>
        <Hero title="Page not found" subtitle="The page you were looking for couldn't be found."/>
        <section className="section">
            <div className="container">
                The page you are looking for couldn't be found.
            </div>
        </section>
    </Page>
}