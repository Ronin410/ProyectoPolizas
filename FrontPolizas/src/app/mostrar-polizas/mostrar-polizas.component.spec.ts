import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MostrarPolizasComponent } from './mostrar-polizas.component';

describe('MostrarPolizasComponent', () => {
  let component: MostrarPolizasComponent;
  let fixture: ComponentFixture<MostrarPolizasComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [MostrarPolizasComponent]
    });
    fixture = TestBed.createComponent(MostrarPolizasComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
