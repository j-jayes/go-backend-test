<script>
    export let name;
    import { onMount } from "svelte";

    let books = [];
    let loading = true;
    let error = null;

    onMount(async () => {
        try {
            const res = await fetch('/books');
            if (!res.ok) {
                throw new Error('Network response was not ok');
            }
            books = await res.json();
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    });
</script>

<main>
    <h1>Hello {name}!</h1>
    <p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
    
    {#if loading}
        <p>Loading...</p>
    {:else if error}
        <p>Error fetching books: {error}</p>
    {:else}
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
                        <td>{book.ID}</td>
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
    {/if}
</main>

<style>
    /* ... Your existing styles ... */
</style>
