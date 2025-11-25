import XCTest
import SwiftTreeSitter
import TreeSitterScaf

final class TreeSitterScafTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_scaf())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Scaf grammar")
    }
}
