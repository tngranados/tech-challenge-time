import { ApiBase } from './Base';
import { Session } from "../models/Session";

export class SessionsApi extends ApiBase {
    getAll(): Promise<Session[]> {
        return this.fetch<Session[]>('sessions', {
            method: 'GET'
        })
    }

    getFinished(): Promise<Session[]> {
        return this.fetch<Session[]>('finished-sessions', {
            method: 'GET'
        })
    }

    getUninished(): Promise<Session[]> {
        return this.fetch<Session[]>('unfinished-sessions', {
            method: 'GET'
        })
    }

    get(id: number) {
        return this.fetch<Session>(`sessions/${id}`, {
            method: 'GET'
        })
    }

    add(session: Session) {
        return this.fetch<string>(`sessions`, {
            method: 'POST',
            body: JSON.stringify(session)
        })
    }

    update(session: Session) {
        return this.fetch<string>(`sessions`, {
            method: 'PUT',
            body: JSON.stringify(session)
        })
    }

    delete(id: number) {
        return this.fetch<string>(`sessions/${id}`, {
            method: 'DELETE'
        })
    }
}
