<script>
	export let name;
	import { onMount } from "svelte";

	let books = [];

	onMount(async () => {
		const res = await fetch('/books');
		books = await res.json();
	});
</script>

<main>
	<h1>Hello {name}!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	<div class="container mt-4">
		<h1 class="mb-4">Book Management</h1>
		<table class="table">
			<thead>
				<tr>
					<th>ID</th>
					<th>Title</th>
					<th>Author</th>
					<th>Quantity</th>
					<th>Actions</th>
				</tr>
			</thead>
			<tbody>
				{#each books as book}
				<tr>
					<td>{book.id}</td>
					<td>{book.title}</td>
					<td>{book.author}</td>
					<td>{book.quantity}</td>
					<td>
						<button class="btn btn-primary mr-2">Checkout</button>
						<button class="btn btn-secondary">Return</button>
					</td>
				</tr>
				{/each}
			</tbody>
		</table>
	</div>
</main>



<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>