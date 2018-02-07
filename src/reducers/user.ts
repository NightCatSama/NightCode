import { UserState, UserAction } from '../model'

const initialState = {
  account: 'nightcat'
}

export default function (state: UserState = initialState, action: UserAction): UserState {
  return state
} 
