# Documentation Standards

## Overview

All code in {{.ProjectName}} must be documented according to these standards. Documentation is not optional - it's a core part of the development process.

## File-Level Documentation

Every source file must start with a file-level comment:

```typescript
/**
 * @fileoverview Brief description of what this file contains
 * 
 * This file provides [main purpose]. It is responsible for:
 * - [responsibility 1]
 * - [responsibility 2]
 * 
 * @module path/to/module
 * @author Project Team
 * @created YYYY-MM-DD
 */
```

```python
"""
Module: module_name

This module provides [main purpose]. It is responsible for:
- [responsibility 1]
- [responsibility 2]

Example:
    from module import function
    result = function(param)
"""
```

## Function Documentation

### TypeScript/JavaScript

```typescript
/**
 * Brief description of what the function does.
 * 
 * @description More detailed description if needed.
 * Can span multiple lines for complex functions.
 * 
 * @param paramName - Description of the parameter
 * @param anotherParam - Description of another parameter
 * @returns Description of what is returned
 * 
 * @throws {ErrorType} When this error occurs
 * 
 * @example
 * ```ts
 * const result = functionName('value', 123);
 * console.log(result); // Expected output
 * ```
 * 
 * @since 1.0.0
 * @see relatedFunction
 */
function functionName(paramName: string, anotherParam: number): ReturnType {
  // implementation
}
```

### Python

```python
def function_name(param_name: str, another_param: int) -> ReturnType:
    """
    Brief description of what the function does.
    
    More detailed description if needed. Can span multiple
    lines for complex functions.
    
    Args:
        param_name: Description of the parameter.
        another_param: Description of another parameter.
    
    Returns:
        Description of what is returned.
    
    Raises:
        ErrorType: When this error occurs.
    
    Example:
        >>> result = function_name("value", 123)
        >>> print(result)
        expected_output
    
    Note:
        Any additional notes about the function.
    """
    # implementation
```

## Class Documentation

```typescript
/**
 * Brief description of the class.
 * 
 * @description More detailed description of the class purpose
 * and how it should be used.
 * 
 * @example
 * ```ts
 * const instance = new ClassName(config);
 * instance.method();
 * ```
 */
class ClassName {
  /**
   * Description of the property.
   */
  public property: Type;

  /**
   * Creates an instance of ClassName.
   * 
   * @param config - Configuration object
   */
  constructor(config: Config) {
    // implementation
  }
}
```

## Component Documentation (React)

```typescript
/**
 * ComponentName - Brief description of what it renders.
 * 
 * @description More detailed description of the component's
 * purpose, behavior, and any important notes.
 * 
 * @example
 * ```tsx
 * <ComponentName 
 *   prop1="value"
 *   prop2={123}
 * />
 * ```
 */
interface ComponentNameProps {
  /** Description of prop1 */
  prop1: string;
  /** Description of prop2 */
  prop2?: number;
}

export function ComponentName({ prop1, prop2 = 0 }: ComponentNameProps) {
  // implementation
}
```

## API Documentation

### Endpoint Documentation

```typescript
/**
 * @route GET /api/users/:id
 * @description Retrieves a user by their ID
 * 
 * @param {string} id - The user's unique identifier
 * 
 * @returns {User} The user object
 * @throws {404} User not found
 * @throws {500} Server error
 * 
 * @example
 * // Request
 * GET /api/users/abc123
 * 
 * // Response
 * {
 *   "id": "abc123",
 *   "name": "John Doe",
 *   "email": "john@example.com"
 * }
 */
```

## README Sections

Every major module should have a README.md with:

1. **Title and Description** - What is this module?
2. **Installation** - How to install/enable it
3. **Usage** - How to use it with examples
4. **API Reference** - Public API documentation
5. **Configuration** - Available configuration options
6. **Examples** - Common use cases
7. **Contributing** - How to contribute (if applicable)

## Inline Comments

Use inline comments for:

1. **Complex logic explanation**: Why this approach was chosen
2. **Non-obvious behavior**: Things that might surprise readers
3. **TODOs**: With ticket/issue reference
4. **Bug workarounds**: With link to issue

```typescript
// Using binary search because the list is sorted
// This gives O(log n) instead of O(n)
const index = binarySearch(sortedList, target);

// TODO(#123): Remove this workaround when API is fixed
// Currently the API returns null for empty strings
if (response.data === null) {
  response.data = '';
}

// HACK: Safari doesn't support this CSS property
// Using vendor prefix as a fallback
element.style.webkitTransform = 'translateX(0)';
```

## Documentation Review Checklist

Before submitting code, verify:

- [ ] All public functions have JSDoc/docstrings
- [ ] Complex logic has inline comments
- [ ] README files are up to date
- [ ] Examples are working and tested
- [ ] Type definitions include descriptions
- [ ] API endpoints are documented
- [ ] Breaking changes are noted
