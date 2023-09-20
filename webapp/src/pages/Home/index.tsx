import preactLogo from '../../assets/preact.svg';
import './style.css';

export function Home() {
	return (
		<div class="home">
			{/* Insert logo once persurl has its own logo / branding */}
			{/*<a href="https://preactjs.com" target="_blank">*/}
			{/*	<img src={preactLogo} alt="Preact logo" height="160" width="160" />*/}
			{/*</a>*/}
			<h1>PersURL - Persistent URL Resolver</h1>
			<h2>This is an evolving project</h2>
			<p>
				Currently PersURL is work-in-progress. Progress is tracked
				on <a href="https://github.com/fabiante/persurl/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc">Github</a>
			</p>
			<p>
				PersURL will expose an HTTP API for integrators which want to create and distribute long-term, persistent
				URLs.
			</p>
			<p>
				The theoretical concept implemented by PersURL are <a href="https://en.wikipedia.org/wiki/Persistent_uniform_resource_locator">PURLs</a>
			</p>
		</div>
	);
}
