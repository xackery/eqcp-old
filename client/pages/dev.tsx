import itemdata from '../items.json' // TESTING

import { ThemeProvider } from '../components/contexts/ThemeContext'
import { HostProvider } from '../components/contexts/HostProvider'
import DataGridView from '../components/data/DataGridView'
import ClassFieldRenderer from '../components/renderers/ClassFieldRenderer'
import ThemeDefinition from '../components/definitions/ThemeDefinition'

const Dev = () => (
  <HostProvider value="http://73.225.105.109:8081/">
  <ThemeProvider value={ThemeDefinition}>
    <div className="container">
      <main>
        <DataGridView fields={['id', {field: 'name', display: "Name"}, {field: 'classes', fieldHandler: ClassFieldRenderer}, 'races', 'lore']} data={itemdata.Items} />
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
  </HostProvider>
)

export default Dev
