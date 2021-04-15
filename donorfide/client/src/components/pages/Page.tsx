import * as React from "react";
import "./Page.sass"

interface pageProps {
    children: React.ReactNode
}

export default function Page(props: pageProps) {
    return <div className="page">
        {props.children}
    </div>
}