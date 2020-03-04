import useSWR from 'swr'

import { ThemeProvider } from '../components/contexts/ThemeContext'
import ThemeDefinition from '../components/definitions/ThemeDefinition'
import ItemHandler from '../handlers/ItemHandler'

const ItemDetails = () => {
  const host = "http://73.225.105.109:8081"
  const handler = new ItemHandler(host)
  const { data, error } = useSWR("1001", handler.getItemById)
  return (
    <ThemeProvider value={ThemeDefinition}>
      <div className="container">
        <main>
          {
            error ? `${error}` :
            !data ? "Loading..." :
            <pre>{JSON.stringify(data, null, 2)}</pre>
          }
        </main>

        <style jsx>{`
          .container {
            min-height: 100vh;
            padding: 0 0.5rem;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            background: ${ThemeDefinition.surface};
            color: ${ThemeDefinition.on_surface};
          }

          main {
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
          }
        `}</style>

        <style jsx global>{`
          html,
          body {
            padding: 0;
            margin: 0;
            background: ${ThemeDefinition.background}
            color: ${ThemeDefinition.on_background}
            font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Oxygen,
              Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue, sans-serif;
          }

          * {
            box-sizing: border-box;
          }
        `}</style>
      </div>
    </ThemeProvider>
  )
}

export default ItemDetails
