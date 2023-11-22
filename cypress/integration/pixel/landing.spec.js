describe('when landing at /...', () => {
  beforeEach(() => {
    cy.visit('http://localhost:3000')
  })

  it('we see static elements', () => {
    cy.log(`sock-image`);
    cy.get(`[data-testid='sock-image']`).should('exist');
    cy.get(`[data-testid='instruction-select-color']`).should('exist');
    // cy.get(`[data-testid='instruction-drop-pixel']`).should('exist');
    // cy.get('.todo-list li').first().should('have.text', 'Pay electric bill')
    // cy.get('.todo-list li').last().should('have.text', 'Walk the dog')
  })
});
