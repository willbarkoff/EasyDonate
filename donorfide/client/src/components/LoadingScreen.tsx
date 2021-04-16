import * as React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleNotch } from "@fortawesome/free-solid-svg-icons";

export default function LoadingScreen(): JSX.Element {
	return <section className="hero is-light is-bold is-fullheight">
		<div className="hero-body">
			<div className="container has-text-centered">
				<FontAwesomeIcon icon={faCircleNotch} spin size="5x" />
			</div>
		</div>
	</section>;
}