import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SuccessDonateComponent } from './success-donate.component';

describe('SuccessDonateComponent', () => {
  let component: SuccessDonateComponent;
  let fixture: ComponentFixture<SuccessDonateComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SuccessDonateComponent]
    });
    fixture = TestBed.createComponent(SuccessDonateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
