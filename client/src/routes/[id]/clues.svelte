<script lang="ts">
    import { page } from "$app/stores";
    import { userId } from "src/store";
    import type { Clue } from "src/types";

    const id = $page.params.id;
    export let conn: WebSocket;

    let movieClue: Clue = {
        word: "movieClue",
        type: "movie",
    };
    let personClue: Clue = {
        word: "personClue",
        type: "person",
    };
    let placeClue: Clue = {
        word: "placeClue",
        type: "place",
    };

    const onSendClues = () => {
        const clues = [movieClue, personClue, placeClue];
        conn.send(
            JSON.stringify({
                type: "send-clues",
                data: JSON.stringify(clues),
                room: id,
                userId: $userId,
            })
        );
    };
</script>

<h1>Clues</h1>

<form on:submit|preventDefault={onSendClues}>
    <label>
        Movie
        <input type="text" bind:value={movieClue.word} />
    </label>
    <label>
        Person
        <input type="text" bind:value={personClue.word} />
    </label>
    <label>
        Place
        <input type="text" bind:value={placeClue.word} />
    </label>
    <button type="submit">Submit</button>
</form>
