function Head () {
    const headerStyle = {
        height: '40px',
        width: '100vw',
        '--wails-draggable': 'drag',
        cursor: 'default'
    }
  return (
      <header style={headerStyle} className={"flex items-center justify-center text-sm"}>
          Conversor Scan Para PDF+OCR
      </header>
  )
}

export default Head;