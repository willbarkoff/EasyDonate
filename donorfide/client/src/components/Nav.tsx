import * as React from "react";
import {OrgContext} from "./App";

export default function Nav() {
    return <OrgContext.Consumer>
        {org => <nav className="navbar" role="navigation" aria-label="main navigation">
            <div className="navbar-brand">
                <a className="navbar-item" href={org.site}>
                    {org.name}
                </a>
            </div>
        </nav>}
    </OrgContext.Consumer>
}