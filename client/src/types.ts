export type User = {
    id: string;
    nickname: string;
    team: string;
    ready: boolean;
    step: number;
};

export type Room = {
    id: string;
    created: string;
    updated: string;
    deleted: string;
    name: string;
    password: string;
}

export type Clue = {
    word: string;
    type: string;
}

export type Round = {
    team: string;
    clue: Clue;
    user: User;
    nextUser: User;
    time: number;
}

