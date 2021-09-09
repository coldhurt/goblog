import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface Article {
  id: string
  title: string
  content: string
}

export interface ArticleState {
  loading: boolean
  articles: Article[]
}

const initialState: ArticleState = {
  loading: false,
  articles: [],
}

export const articleSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    startLoading(state) {
      state.loading = true
    },
    setArticles(state, action: PayloadAction<{ articles: Article[] }>) {
      state.loading = false
      state.articles = action.payload.articles
    },
  },
})

// Action creators are generated for each case reducer function
export const { startLoading, setArticles } = articleSlice.actions

export default articleSlice.reducer
