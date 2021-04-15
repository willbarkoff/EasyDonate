import * as React from 'react'

interface heroProps {
	title?: string
	subtitle?: string
	style?: string
	center?: boolean
}

export default function Hero(props: heroProps) {
	return <div>
		<section className={`hero is-${props.style || "info"} is-bold`}>
			<div className="hero-body">
				<div className={`container ${props.center ? "has-text-centered" : ""}`}>
					<h1 className="title">{props.title}</h1>
					<h2 className="subtitle">{props.subtitle}</h2>
				</div>
			</div>
		</section>
	</div>
}