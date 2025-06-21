/**
 * TemplUI Form Validation - Vanilla JS
 * Similar to React Hook Form but for vanilla JS and templ components
 */

class TemplForm {
    constructor(formElement, options = {}) {
        this.form = formElement;
        this.fields = new Map();
        this.errors = new Map();
        this.touched = new Set();
        this.options = {
            validateOnBlur: true,
            validateOnChange: false,
            validateOnInput: true,
            clearErrorsOnFocus: true,
            clearErrorsOnInput: true,
            showSuccessState: true,
            ...options
        };
        
        this.init();
    }
    
    init() {
        // Find all form fields
        const inputs = this.form.querySelectorAll('input, textarea, select');
        
        inputs.forEach(input => {
            const name = input.name;
            if (!name) return;
            
            this.fields.set(name, {
                element: input,
                rules: this.parseValidationRules(input)
            });
            
            // Add event listeners
            if (this.options.validateOnBlur) {
                input.addEventListener('blur', () => this.validateField(name, 'blur'));
            }
            
            if (this.options.validateOnChange) {
                input.addEventListener('change', () => this.validateField(name, 'change'));
            }
            
            if (this.options.validateOnInput) {
                input.addEventListener('input', () => {
                    if (this.touched.has(name) || this.errors.has(name)) {
                        this.validateField(name, 'input');
                    }
                });
            }
            
            // Clear errors on focus
            if (this.options.clearErrorsOnFocus) {
                input.addEventListener('focus', () => {
                    if (this.errors.has(name)) {
                        this.clearFieldError(name);
                    }
                });
            }
            
            // Clear errors on input (immediate feedback)
            if (this.options.clearErrorsOnInput) {
                input.addEventListener('input', () => {
                    if (this.errors.has(name) && input.value.trim()) {
                        // Give immediate positive feedback by clearing error
                        setTimeout(() => {
                            if (!this.errors.has(name)) return;
                            this.validateField(name, 'input-clear');
                        }, 100);
                    }
                });
            }
        });
        
        // Handle form submission
        this.form.addEventListener('submit', (e) => {
            e.preventDefault();
            this.handleSubmit();
        });
    }
    
    parseValidationRules(input) {
        const rules = [];
        
        // Required
        if (input.hasAttribute('required') || input.dataset.required === 'true') {
            rules.push({
                type: 'required',
                message: 'This field is required'
            });
        }
        
        // Min length
        if (input.dataset.min) {
            rules.push({
                type: 'min',
                value: parseInt(input.dataset.min),
                message: `Minimum ${input.dataset.min} characters required`
            });
        }
        
        // Max length
        if (input.dataset.max) {
            rules.push({
                type: 'max',
                value: parseInt(input.dataset.max),
                message: `Maximum ${input.dataset.max} characters allowed`
            });
        }
        
        // Pattern
        if (input.dataset.pattern) {
            rules.push({
                type: 'pattern',
                value: new RegExp(input.dataset.pattern),
                message: input.dataset.patternMessage || 'Invalid format'
            });
        }
        
        // Email
        if (input.type === 'email') {
            rules.push({
                type: 'email',
                message: 'Please enter a valid email address'
            });
        }
        
        return rules;
    }
    
    validateField(fieldName, trigger = null) {
        const field = this.fields.get(fieldName);
        if (!field) return true;
        
        const value = field.element.value.trim();
        const errors = [];
        
        // Mark as touched on certain triggers
        if (trigger === 'blur' || trigger === 'change') {
            this.touched.add(fieldName);
        }
        
        // Skip validation on focus clear
        if (trigger === 'focus-clear') {
            return true;
        }
        
        // Validate each rule
        for (const rule of field.rules) {
            const error = this.validateRule(value, rule);
            if (error) {
                errors.push(error);
                break; // Stop at first error
            }
        }
        
        if (errors.length > 0) {
            this.setFieldError(fieldName, errors[0], trigger);
            return false;
        } else {
            this.clearFieldError(fieldName, trigger);
            return true;
        }
    }
    
    validateRule(value, rule) {
        switch (rule.type) {
            case 'required':
                return !value ? rule.message : null;
                
            case 'min':
                return value.length < rule.value ? rule.message : null;
                
            case 'max':
                return value.length > rule.value ? rule.message : null;
                
            case 'pattern':
                return !rule.value.test(value) ? rule.message : null;
                
            case 'email':
                const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
                return !emailRegex.test(value) ? rule.message : null;
                
            default:
                return null;
        }
    }
    
    setFieldError(fieldName, message, trigger = null) {
        this.errors.set(fieldName, message);
        
        const field = this.fields.get(fieldName);
        if (!field) return;
        
        // Add error styling to input with animation
        field.element.classList.add('border-red-500', 'focus:border-red-500');
        field.element.classList.remove('border-green-500', 'border-yellow-500');
        
        // Show error message in form.Message() component
        const messageElement = this.form.querySelector(`[data-field="${fieldName}"]`);
        if (messageElement) {
            messageElement.textContent = message;
            // Keep existing classes but ensure error styling
            const existingClasses = messageElement.className;
            if (!existingClasses.includes('text-red-500')) {
                messageElement.className = existingClasses.replace('text-blue-500', 'text-red-500');
            }
            messageElement.style.display = 'block';
        }
        
        // Add shake animation on blur validation
        if (trigger === 'blur') {
            field.element.classList.add('animate-pulse');
            setTimeout(() => field.element.classList.remove('animate-pulse'), 600);
        }
    }
    
    clearFieldError(fieldName, trigger = null) {
        this.errors.delete(fieldName);
        
        const field = this.fields.get(fieldName);
        if (!field) return;
        
        // Remove error styling
        field.element.classList.remove('border-red-500', 'focus:border-red-500');
        
        // Add success styling if field has value and success state is enabled
        if (this.options.showSuccessState && field.element.value.trim() && trigger !== 'focus-clear') {
            field.element.classList.add('border-green-500');
            // Remove success styling after a delay if user continues typing
            if (trigger === 'input' || trigger === 'input-clear') {
                setTimeout(() => {
                    if (!this.errors.has(fieldName)) {
                        field.element.classList.remove('border-green-500');
                    }
                }, 2000);
            }
        }
        
        // Hide error message
        const messageElement = this.form.querySelector(`[data-field="${fieldName}"]`);
        if (messageElement) {
            messageElement.style.display = 'none';
        }
    }
    
    validateAll() {
        let isValid = true;
        
        for (const [fieldName] of this.fields) {
            if (!this.validateField(fieldName)) {
                isValid = false;
            }
        }
        
        return isValid;
    }
    
    getValues() {
        const values = {};
        
        for (const [fieldName, field] of this.fields) {
            values[fieldName] = field.element.value;
        }
        
        return values;
    }
    
    handleSubmit() {
        if (this.validateAll()) {
            const values = this.getValues();
            
            // Trigger custom submit event
            const submitEvent = new CustomEvent('formSubmit', {
                detail: { values, form: this }
            });
            
            this.form.dispatchEvent(submitEvent);
        }
    }
    
    // Public API
    reset() {
        this.form.reset();
        this.errors.clear();
        this.touched.clear();
        
        // Clear all error messages and styling
        for (const [fieldName, field] of this.fields) {
            field.element.classList.remove('border-red-500', 'focus:border-red-500', 'border-green-500');
            this.clearFieldError(fieldName, 'reset');
        }
    }
    
    // Configure validation triggers
    setValidationTriggers(options) {
        this.options = { ...this.options, ...options };
        // Re-initialize with new options
        this.init();
    }
    
    // Manually trigger validation for a field
    validateFieldManually(fieldName) {
        return this.validateField(fieldName, 'manual');
    }
    
    // Set custom validation rule for a field
    setCustomRule(fieldName, rule) {
        const field = this.fields.get(fieldName);
        if (field) {
            field.rules.push(rule);
        }
    }
    
    getErrors() {
        return Object.fromEntries(this.errors);
    }
    
    hasErrors() {
        return this.errors.size > 0;
    }
}

// Auto-initialize forms
document.addEventListener('DOMContentLoaded', () => {
    const forms = document.querySelectorAll('[data-form-validate="true"]');
    
    forms.forEach(form => {
        // Read validation options from data attributes
        const options = {
            validateOnBlur: form.dataset.validateOnBlur !== 'false',
            validateOnChange: form.dataset.validateOnChange === 'true',
            validateOnInput: form.dataset.validateOnInput !== 'false',
            clearErrorsOnFocus: form.dataset.clearErrorsOnFocus !== 'false',
            clearErrorsOnInput: form.dataset.clearErrorsOnInput !== 'false',
            showSuccessState: form.dataset.showSuccessState !== 'false'
        };
        
        const templForm = new TemplForm(form, options);
        
        // Store instance on form element for external access
        form._templForm = templForm;
        
        // Handle form submission
        form.addEventListener('formSubmit', (e) => {
            const { values } = e.detail;
            console.log('Form submitted:', values);
            
            // Here you would typically send data to server
            // fetch(form.action, { method: form.method, body: new FormData(form) })
        });
    });
});

// Export for module usage
if (typeof module !== 'undefined' && module.exports) {
    module.exports = TemplForm;
}
