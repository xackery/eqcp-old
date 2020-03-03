import { useContext } from 'react';
import ThemeContext from '../contexts/ThemeContext'

const DataGridView = ({fields, data}) => {
  const theme = useContext(ThemeContext)
  return (
    <div>
      <table>
        <thead>
          <tr>
            {
              fields.map((value, index) => {
                if (typeof(value) === "string") {
                  return (
                    <th key={index}>{value}</th>
                  )
                } else {
                  return (
                    <th key={index}>{value.display || value.field}</th>
                  )
                }
              })
            }
          </tr>
        </thead>
        <tbody>
          {
            data.map((value) => {
              return (
                <tr key={value.id}>
                  {
                    fields.map((field) => {
                      if (typeof(field) === "string") {
                        return (
                          <td key={field}>
                            {value[field]}
                          </td>
                        )
                      } else {
                        return (
                          <td key={field.field}>
                            { field.fieldHandler ?
                              <field.fieldHandler value={value[field.field]} /> :
                              <span>{value[field.field]}</span>
                            }
                          </td>
                        )
                      }
                    })
                  }
                </tr>
              )
            })
          }
        </tbody>
      </table>

      <style jsx>{`
          table {
            border-collapse: collapse;
            border-radius: 5px 5px 0 0;
            overflow: hidden;
            margin: 25px 0;
            font-size: 0.9em;
            min-width: 400px;
            background: ${theme.element};
            color: ${theme.on_element};
            box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);
          }
          table thead tr {
            background: ${theme.primary};
            color: ${theme.on_primary};
            text-align: left;
          }

          table th,
          table td {
            padding: 12px 15px;
          }

          table tr {
            border-bottom: 1px solid ${theme.surface}
          }

          table tbody tr:nth-of-type(even) {
            background: ${theme.element_raised}
          }

          table tbody tr:last-of-type {
            border-bottom: 2px solid ${theme.primary};
          }
      `}</style>
    </div>
  )
}

export default DataGridView