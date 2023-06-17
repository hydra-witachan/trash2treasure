import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ConfirmationTopUpComponent } from './confirmation-top-up.component';

describe('ConfirmationTopUpComponent', () => {
  let component: ConfirmationTopUpComponent;
  let fixture: ComponentFixture<ConfirmationTopUpComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ConfirmationTopUpComponent]
    });
    fixture = TestBed.createComponent(ConfirmationTopUpComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
