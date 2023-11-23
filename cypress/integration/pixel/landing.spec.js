describe('when landing at /...', () => {
  beforeEach(() => {
    cy.visit('http://localhost:3000')
  })

  it('we see static elements', () => {
    cy.log(`sock-image`);
    cy.getByTestId('title').should('exist').contains(`pixel-socks`);
    cy.get(`[data-testid='sock-image']`).should('exist');
    cy.get(`[data-testid='instruction-select-color']`).should('exist').contains(`select a color`);
    cy.get(`[data-testid='color-picker']`).should('exist');
    cy.get(`[data-testid='color-picker']`).click();
  })
  it('color picker opens on click and allows color pick', () => {
    cy.get(`[data-testid='color-picker']`).click();
    cy.get(`.ant-color-picker-saturation`).should('exist').should('be.visible');
    cy.get(`.ant-color-picker-palette`).click({multiple: true});
    cy.wait(2000);
    cy.getByTestId(`main`).should('exist').click(50, 50);
    cy.wait(2000);
    cy.getByTestId(`color-r`).should('exist').invoke('text').then((value) => {
      try {
        parseInt(value);
        cy.log(`found color component: ${value}`);
      } catch (err) {
        cy.log(`color value must be numeric, instead found ${value}`);
        throw err
      }
    });
    cy.getByTestId(`color-g`).should('exist').invoke('text').then((value) => {
      try {
        parseInt(value);
        cy.log(`found color component: ${value}`);
      } catch (err) {
        cy.log(`color value must be numeric, instead found ${value}`);
        throw err
      }
    });
    cy.getByTestId(`color-b`).should('exist').invoke('text').then((value) => {
      try {
        parseInt(value);
        cy.log(`found color component: ${value}`);
      } catch (err) {
        cy.log(`color value must be numeric, instead found ${value}`);
        throw err
      }
    });
    cy.getByTestId('instruction-drop-pixel').should('exist').should('be.visible').contains('drop a pixel');
  })
});
