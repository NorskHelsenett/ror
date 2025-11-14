# Architecture Decision Record: React framework for ROR Web GUI

## Summary

### Issue
We want to use a React framework to create our web application:
  * We want to use React because it is easier to find people who know the framework
  * We want to use a meta-framework that enables to do more heavy-lifting with simple configuration or conventions

### Decision
Decided on React Router v7. Decision made by @patrickedqvist, @havardelnan and @rogerwesterbo.

### Status
Decided on React Router v7. If we are hindered by our choice the backup choice is Next.js.

## Details
Two choices needed to be made. One is which framework to choose and the other choice which meta-framework should be used. The first choice has already been decided to be React due it's much higher usage [^1]. It will simply be easier to acquire talent for a framework many people know. That leaves choosing a meta-framework or some type of tooling the meet our needs listed below.

### Needs
- Support for running on node but preferably either on Bun or Deno
- Server-Side Rendering (SSR) out-of-the-box or with minimal configuration
- Routing
- Code-splitting capabilities
- Support for typescript
### Nice to haves
- Configurable cache-mechanics
- Serverless support (not sure how to utilities it properly, discuss.)
- Support for telemetry
### Not of particular interest
- Small client-side JS bundle size
- Can be dependent on JS to exist on the client machine
## Meta frameworks
These are some known meta frameworks and toolings that support React that we will compare to our needs and wishes. We also compare some features to highlight how the different frameworks has solved them, such as how data-fetching is done.

1. Next.js
2. Astro
3. React Router v7

Remix is not considered as React Router v7 has replaced it (https://remix.run/blog/incremental-path-to-react-19). Vite with SSR is not considered since it is a low-level API for framework developers, but React Router v7 is built on this technology. Rakkas.js is another framework that was on the radar but it is not production-ready at this moment.

| Feature                          | Next.js                                              | Astro                                | React Router v7                            |
| -------------------------------- | ---------------------------------------------------- | ------------------------------------ | ------------------------------------------ |
| Supports Node.js                 | Yes                                                  | Yes                                  | Yes                                        |
| Supports Bun and/or Deno runtime | Yes                                                  | Yes                                  | Unknown                                    |
| Server-side rendering            | Yes                                                  | Yes                                  | Yes [^3]                                   |
| Routing                          | File-based                                           | File-based                           | Config-based                               |
| Code-Splitting                   | By route                                             | Granular (Island-Architecture)       | By route                                   |
| Configurable cache-mechanics     | Yes [^2]                                             | No                                   | Yes [^4]                                   |
| Methods for data-fetching        | Standard async/await in with React Server Components | Standard async/await in .astro files | clientLoader, loader, clientAction, action |
| POST request handling            | Yes                                                  | No                                   | Yes                                        |
| Middleware                       | Yes                                                  | Yes                                  | Yes                                        |
| Telemetry                        | Yes                                                  | No                                   | No                                         |
| Support for Typescript           | Yes                                                  | Yes                                  | Yes                                        |

## Framework differences

While all these frameworks has a similar set of features there are architectural difference, they might have different levels of complexity for good and bad. In this section I try to highlight how the same set of features are solved differently.

### Server-side rendering (SSR) and data-fetching

#### Next.js
Next.js uses [React Server Components](https://react.dev/reference/rsc/server-components) per default. It enables the use of async and await inside if your component without any extra feature. You can also opt-out and use client-side data fetching for the places you want.

**Example**
```
import db from './database';

async function Note({id}) {
	// NOTE: loads *during* render.
	const note = await db.notes.get(id);
	return (
		<div>
			<Author id={note.authorId} />
			<p>{note}</p>
		</div>
	);
}

async function Author({id}) {
	// NOTE: loads *after* Note,
	// but is fast if data is co-located.
	const author = await db.authors.get(id);
	return <span>By: {author.name}</span>;
}
```

#### Astro
The framework for content-driven websites. Astro has the possibility for using SSR but its speciality is generating static content. It follows an island-architecture approach where you create the main content as static but can create "islands" or "regions" where you can use one of the major client-side libraries (react, vue, svelte etc) to render an isolated interactive experience.

> Islands architecture works by rendering the majority of your page to fast, static HTML with smaller “islands” of JavaScript added when interactivity or personalization is needed on the page (an image carousel, for example). [^6]

SSR is available but needs a runtime adapter to act as a server. Currently there are adapters for the popular JAM-stack hosting partners such as vercel and netlify. There is also a node adapter.

**Example**
```
---
export const prerender = false; // Not needed in configured 'server' mode

import { getProduct } from '../api';

const product = await getProduct(Astro.params.id);

// No product found
if (!product) {
	return new Response(null, {
		status: 404,
		statusText: 'Not found'
	});
}

// The product is no longer available
if (!product.isAvailable) {
	return Astro.redirect("/products", 301);
}

---

<html>

<!-- Page here... -->

</html>
```

#### React Router v7
Version 7 of React Router is the next iteration of Remix. Remix was the best alternative to Next.js offering much of the same features. React Router uses some different namespaced methods for doing server-side fetching and rendering. In the example below it uses the function "loader" to fetch the data on the server. There are then functions for doing posting (usually used together with a form) and both fetching and posting on the client side.

**Example**
```
// route("products/:pid", "./product.tsx");
import type { Route } from "./+types/product";
import { fakeDb } from "../db";

export async function loader({ params }: Route.LoaderArgs) {
  const product = await fakeDb.getProduct(params.pid);
  return product;
}

export default function Product({
  loaderData,
}: Route.ComponentProps) {
  const { name, description } = loaderData;
  return (
    <div>
      <h1>{name}</h1>
      <p>{description}</p>
    </div>
  );
}
```

The other big difference from the other frameworks is that uses a configuration for the routing rather than a file-based approach, although there is a possibility to implement file-based routing as well or even mix the two.

**Example**
```
import {
  type RouteConfig,
  route,
  index,
  layout,
  prefix,
} from "@react-router/dev/routes";

export default [
  index("./home.tsx"),
  route("about", "./about.tsx"),

  layout("./auth/layout.tsx", [
    route("login", "./auth/login.tsx"),
    route("register", "./auth/register.tsx"),
  ]),

  ...prefix("concerts", [
    index("./concerts/home.tsx"),
    route(":city", "./concerts/city.tsx"),
    route("trending", "./concerts/trending.tsx"),
  ]),
] satisfies RouteConfig;
```

### ROR considerations

Since Astro is primarily meant for content-driven websites I think it's safe to assume that what we want to build diverge from their intention and ambition. We are looking for a framework that best fits building a data-heavy dashboard application. So I will simply write my pros and cons for Next.js and React Router v7.

#### Next.js Pros
- Leans heavy on React v19 API and the two frameworks are constantly in dialog pushing React forward. If you learn and keep updated on React there are less boilerplate to think about for Next.js
- The most popular option - easy to find compatible libraries and people who have experience with it.
- Mature framework means flexibility and extendability without offering stability. Next.js offers many different features that are battle-proven.
#### Next.js Cons
- Opiniated, it for instance builds on extending certain built-in standards. There for there are somethings that are Next.js specific that there are minimal documentation on.
- Vendor lock-in. Vercel is the company behind Next.js which mainly earns money from their hosting solution. Next.js and Vercels hosting solution is tightly coupled meaning that other alternatives does make full use out of the capabilities.

#### React Router Pros
- Very flexible routing could allow for an easier file-organization
- Familiarity - many people who have worked with React Single-Page-Applications have used React Router v6 or older versions and therefore there might be less of a hurdle to work with the framework.
- React Router uses the native Request and Response API making it easier for everyone to rely on MDN for documentation.
- Uses Vite under the hood which means it can use other peoples work to grow or for us to modify certain behaviour.

#### React Router Cons
- Excessive complexity, one of the pain points for many developers when it comes to meta-frameworks. React Router uses some of their own functionality for data-fetching, inherited from Remix. Since React 19 there is the question if these are really necessary or if the framework should solely rely on what is possible with for instance server actions and functions.
- React-Router in its current state does not extend much beyond routing and data-fetching. Meaning certain features can be harder to implement since it means configuring or extending Vite.



[^1]:  Source https://share.stateofjs.com/share/prerendered?localeId=en-US&surveyId=state_of_js&editionId=js2024&blockId=front_end_frameworks_ratios&params=&sectionId=libraries&subSectionId=front_end_frameworks ![[Front End Frameworks Ratios.png]]

[^2]: https://nextjs.org/docs/app/building-your-application/caching
[^3]: They say it is supported but "Server side rendering requires a deployment that supports it." without any further explanation. https://reactrouter.com/start/framework/rendering. Probably means you have setup your own server without guidance from the framework.
[^4]: By using custom Cache-Control headers on each route, note that it is not granular for different type of requests.
[^5]: SolidStart is built on top of https://nitro.build/ which has granular cache control.
[^6]: https://docs.astro.build/en/concepts/islands/
