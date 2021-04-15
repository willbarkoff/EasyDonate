import * as React from "react";
import {OrgContext} from "./App";

export default function Footer() {
    return <OrgContext.Consumer>
        {org => <footer className="footer">
            <div className="container has-text-centered">
                <p>{org.imprint}</p>
                <p>Powered by <a href="https://donorfide.org" rel="noopener noreferrer" target="_blank">Donorfide</a>.
                </p>
            </div>
        </footer>
        }
    </OrgContext.Consumer>
}
